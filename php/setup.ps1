# start the server in the background and give it a second to attach to the network interface
$job = php ./src/server.php &
sleep 1

# kill the server when this script completes
trap {$job | Remove-Job}

php ./tests/test.php