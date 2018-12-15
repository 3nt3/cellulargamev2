var ws = new WebSocket("ws://localhost:8000/");

ws.onmessage = function() {
	console.log(event.data);
}

function spawnFood() {
	ws.send(JSON.stringify({type: "spawnFood"}))
}