require 'socket'

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
          message = socket.readline.strip
          socket.puts(message)
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
