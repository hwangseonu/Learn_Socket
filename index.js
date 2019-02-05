'use strict';

const express = require('express');
const http = require('http');
const socket = require('socket.io');

const app = express();
const server  = http.Server(app);
const io = socket(server);

app.use(express.static('public'));

io.on('connection', function (socket) {
  socket.on('chat message', function (msg) {
    io.emit('chat message', msg);
  })
});

const port = 3000;
server.listen(port, () => console.log(`server is running on port ${port}`));
