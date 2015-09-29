mysql = require 'mysql'
fs = require "fs"
_ = require 'underscore'
S = require 'underscore.string'
ejs = require 'ejs'

connection =
  mysql.createConnection
    host: 'localhost'
    user: 'humhub'
    password: 'humhub'
    database: 'humhub'

connection.connect()

query = (sql, callback) ->
  connection.query sql, (err, rows, fields) ->
    if err
      console.log err
    else
      callback(rows, fields)

listTables = ->
  query 'show tables', (rows, fields) ->
    tables = _(rows).map( (r) -> r['Tables_in_humhub'])
    _(tables).forEach( (t) -> writeTableStruct t)

###
listTableColumns = (table) ->
  query "describe #{table}", (rows, fields) ->
    console.log "TABLE: #{table}"
    _(rows).forEach (r) ->
      attribute = S(r.Field).classify().capitalize().value()
      console.log "#{attribute} -> #{r.Type}"
      connection.end()
###

class Attribute
  constructor: (@name, @type, @tag) ->

class Table
  constructor: (@name) ->
    @trace = false
    @attributes = []
  ###
  constructor: ->
    @trace: false
    @attributes: []
    @types: []
    @jsonAttribute: []
  ###
  addAttribtue: (name, type, tag) ->
    newAttribute = new Attribute(name, type, tag)
    @attributes.push(newAttribute)

modifyFieldType = (type) ->
    if type.indexOf("int") != -1  # type contains (int, tinyint)
      return "uint"
    else if type.indexOf("time") != -1  #type contains (date_time)
      return "time.Time"
    else  #type contains (text, varchar)
      return "string"


traceTable = (attribute) ->
  switch attribute
    when "CreatedAt", "CreatedBy", "UpdatedAt", "UpdatedBy"
      return true
    else
      return false

###
writeIntoFile = (data) ->
  fs.appendFile 'models.go', data + "\n", (error) ->
    console.error("Error writing file", error) if error
###

writeTableStruct = (tableName) ->

  fs.writeFile './models/models.go', ''

  query "describe humhub.#{tableName}", (rows, fields) ->
    table = new Table(S(tableName).classify().capitalize().value())

    ###
    # write tableId struct
    console.log "type #{t.name}Id struct {"
    console.log "    Id uint `json:\"id\" sql:\"AUTO_INCREMENT\"`"
    console.log "}\n"
    ###

    ###
    writeIntoFile("type #{tableName}Id struct {")
    writeIntoFile("    Id uint `json:\"id\" sql:\"AUTO_INCREMENT\"`")
    writeIntoFile("}")
    ###

    # console.log "type #{t.name}Data struct {"
    # write tableData struct

    #fs.writeFile './models.go', '', (err) ->
    #  console.log 'error in writing file'

    _(rows).forEach (r) ->
      attribute = S(r.Field).classify().capitalize().value()
      type = r.Type
      if attribute != "Id"
        if traceTable(attribute)
          table.trace = true
        else
          modelType = modifyFieldType(type)  # convert sqlType to modelType
          jsonAttribute = attribute[0].toLowerCase() + attribute.substr(1)  # lowercase the attribute for json
          table.addAttribtue(attribute, modelType, jsonAttribute)

    file = fs.readFileSync './modelTemplate.ejs', 'utf8'
    data = ejs.render(file, table)

    fs.appendFile './models/models.go', data

    #fs.appendFile 'model.go', data, (error) ->
    #  console.error("Error writing file", error) if error
        # console.log "    #{attribute} #{modelType} `json:\"#{jsonAttribute}\"`"
    # console.log "}\n"

    ###
    # write table struct
    console.log "type #{tableName} struct {"
    console.log "#{tableName}Id"
    console.log "#{tableName}Data `json:\"data\"`"
    if t.trace
      console.log "Trace `json:\"trace\"`"
    console.log "}\n"
    ###

    ###
    console.log "type #{table.name} struct {"
    for i in [0...table.attributes.length]
      console.log "   #{table.attributes[i]} #{table.types[i]} `json:\"#{table.jsonAttribute[i]}\"`"
    console.log "Trace: #{table.trace}"
    console.log "}"
    ###

listTables()
