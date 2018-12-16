var ws = new WebSocket("ws://localhost:8000/");

ws.onmessage = function() {
	console.log(JSON.parse(atob(JSON.parse(event.data).data)));
}

function spawnFood() {
	ws.send(JSON.stringify({type: "spawnFood"}))
}

function getFood() {
	ws.send(JSON.stringify({type: "getFood"}))
}

function initCell(name) {
	ws.send(JSON.stringify({type: "initCell", data: JSON.stringify({name: name})}))
}

function getCells() {
	ws.send(JSON.stringify({type: "getCells"}))
}

function eat(id, mealId) {
	ws.send(JSON.stringify({type: "eat", data: JSON.stringify({id: id, mealId: mealId})}));
}

function updtaeSize(id, size) {
	ws.send(JSON.stringify({type: "updateSize", data: JSON.stringify({id: id, size: size})}));
}

function changePos(id, pos) {
	ws.send(JSON.stringify({type: "changePos", data: JSON.stringify({id: id, posX: pos[0], posY: pos[1]})}));
}
