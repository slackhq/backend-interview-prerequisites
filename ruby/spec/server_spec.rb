require_relative '../lib/server'

describe Server do
  before(:all) do
    @server = Server.new
    @server.listen
  end

  after(:all) do
    @server.shutdown
  end

  it 'pings' do
    socket = TCPSocket.open('localhost', Server::DEFAULT_PORT)
    socket.puts('ping')
    response = socket.readline.strip
    expect(response).to eq('ping')
  end
end
