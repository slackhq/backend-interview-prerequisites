require 'socket'

require_relative 'message'

class Server
  DEFAULT_PORT = 8000

  def initialize(port = DEFAULT_PORT)
    @server = TCPServer.new(port)
    @stop = false
  end

  def listen
    Thread.new do
      while (!stop && s = server.accept)
        Thread.new(s) do |socket|
          text = socket.readline.strip
          Message.create(msg: text)
          m = Message.last
          socket.puts(m.msg)
          socket.close
        end
      end
    end
  end

  def shutdown
    @stop = true
  end

  private

  attr_reader :server, :stop
end
