pg  = require 'pg'
fs  = require "fs"
_   = require 'underscore'
S   = require 'underscore.string'
ejs = require 'ejs'

host     = 'localhost'
username = 'admin'
password = 'admin'
database = 'nxt'
schema   = 'lampetia'
tables   = ['event_ticket','event','agenda_slot','agenda_row','agenda',
           'speaker','session_speaker','agenda_day','session']
pg_conn_str = "postgres://#{username}:#{password}@#{host}/#{database}"
pg_client   = new pg.Client pg_conn_str

# ------------------------------------------------------------------------------
init_models = ->
  fs.exists __dirname+'/model', (model_dir_exists) ->
    unless model_dir_exists
      fs.mkdir __dirname+'/model'
    file = fs.readFileSync './init_model.ejs', 'utf8'
    data = ejs.render(file)
    fs.writeFileSync './model/models.go', data
    generate_models()

generate_models = ()->
  pg_client.connect (err)->
    if err
      console.error "couldn't connect to postgres", err
    else
      tables.forEach (table_name)->
        generate_table_model table_name

      pg_client.end
# ------------------------------------------------------------------------------
get_col_obj = (col)->
  col_obj =
    id   : prepare_identifier(col['column_name'])
    type : get_go_type(col["data_type"])
    tag  : col['column_name']
  return col_obj
# ------------------------------------------------------------------------------
generate_table_model = (table_name)->
  table_obj = new Table(prepare_identifier(table_name))
  query_cols_inf = "SELECT *
    FROM information_schema.columns
    WHERE table_schema = '#{schema}'
    AND table_name   = '#{table_name}'"
  pg_client.query query_cols_inf, (err, res) ->
    columns = []
    for r in res.rows
      col = get_col_obj r
      table_obj.addAttribtue(col.id, col.type,col.tag)

    file = fs.readFileSync './model_template.ejs', 'utf8'
    data = ejs.render(file, table_obj)
    fs.appendFile './model/models.go', data

# ------------------------------------------------------------------------------
get_go_type = (pg_type)->
  if pg_type.indexOf("int") != -1 or pg_type.indexOf("numeric") != -1
    return "uint"
  else if pg_type.indexOf("time") != -1
    return "time.Time"
  else if pg_type.indexOf("double") != -1
    return "float64"
  else  #type contains (text, varchar)
    return "string"
# ------------------------------------------------------------------------------
class Attribute
  constructor: (@name, @type, @tag) ->

class Table
  constructor: (@name) ->
    @trace = false
    @attributes = []

  addAttribtue: (name, type, tag) ->
    newAttribute = new Attribute(name, type, tag)
    @attributes.push(newAttribute)

prepare_identifier = (id)->
  return S(id).classify().capitalize().value()
# ------------------------------------------------------------------------------
init_models()
