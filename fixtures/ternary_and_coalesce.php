<?php

$a ? $b : $c;
$a ?: $c;

$a ? $b : $c ? $d : $e;
$a ? $b : ($c ? $d : $e);

$a ?? $b;
$a ?? $b ?? $c;
$a ?? $b ? $c : $d;
$a && $b ?? $c;
