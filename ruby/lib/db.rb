require 'active_record'
require 'sqlite3'
require 'fileutils'
require 'tempfile'

module DB
  def self.init
    @datastore = Tempfile.new
    create_schema
    @datastore.close
    connect
  end

  def self.cleanup
    @datastore.unlink
  end

  private

  def self.create_schema
    db = SQLite3::Database.open(@datastore.path)
    db.execute("CREATE TABLE messages (id INTEGER NOT NULL PRIMARY KEY, msg TEXT NOT NULL);")
  end

  def self.connect
    ActiveRecord::Base.establish_connection(
        adapter: :sqlite3,
        database: @datastore.path
    )
  end
end
