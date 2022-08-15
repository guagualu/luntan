//加载express模块
var express = require('express');
var app = express();

app.use('/storage/uploads', express.static('storage/uploads'));

var server = app.listen(8081);
