// Chat.js
import React, { useState } from "react";

import io from "socket.io-client";
import useChatStore from "./stores/chat.store";

const Chat = () => {
	const { messages, socket } = useChatStore();
	const [input, setInput] = useState("");

	const sendMessage = () => {
		if (socket && input.trim()) {
			// Emit a message event to the server
			socket.emit("message", input);
			setInput("");
		}
	};

	return (
		<div>
			<div>
				<h2>Messages</h2>
				<ul>
					{messages.map((msg, idx) => (
						<li key={idx}>{msg}</li>
					))}
				</ul>
			</div>
			<input
				type="text"
				value={input}
				onChange={(e) => setInput(e.target.value)}
				placeholder="Type your message..."
			/>
			<button onClick={sendMessage}>Send</button>
		</div>
	);
};

export default Chat;
