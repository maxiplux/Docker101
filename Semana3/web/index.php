<?php
header('Content-Type: text/plain', true, 200);
echo "Hello, DevHack!!\n";
$redis = new Redis();
$redis->connect('redis');
echo 'Number of visits: ' . $redis->incr('visitors') . "\n";