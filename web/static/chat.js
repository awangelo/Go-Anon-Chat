const ws = new WebSocket("ws://localhost:8080/chat");

ws.onopen = function (event) {
    console.log("Conectado ao servidor WebSocket");
};

ws.onmessage = function (event) {
    const messagesDiv = document.getElementById("messages");
    const message = document.createElement("div");
    message.innerHTML = event.data;

    // Verificar se a mensagem eh a contagem de usuários
    const userCountDiv = message.querySelector("#user-count");
    if (userCountDiv) {
        document.getElementById("user-count").innerHTML = userCountDiv.innerHTML;
    } else {
        messagesDiv.appendChild(message);
        messagesDiv.scrollTop = messagesDiv.scrollHeight;
    }
};

ws.onclose = function (event) {
    console.log("Desconectado do servidor WebSocket");
};

function sendMessage() {
    const input = document.getElementById("messageInput");
    const message = input.value;
    if (message) {
        ws.send(message);
        input.value = "";
    }
}

// Enviar com Enter.
document.getElementById("messageInput").addEventListener("keydown", function(event) {
    if (event.key === "Enter") {
        sendMessage();
    }
});