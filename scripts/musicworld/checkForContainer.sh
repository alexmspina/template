if [[ -n "docker ps -q -f name=musicworld_dev"]]; then

    echo "found musicworld_dev"

    if [[ -n "docker ps -aq -f status=exited -f name=musicworld_dev" ]]; then
    	echo "musicworld_dev exited"
	docker ps -aq -f status=exited -f name=musicworld_dev
    fi
fi
if [[ -z "docker ps -q -f name=musicworld_dev" ]]; then

    echo "could not find musicworld_dev"
fi
