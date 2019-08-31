const mix = require('laravel-mix');
require('mix-env-file');

mix.env('.env');

mix.react('client_src/app_main.js', 'static/app.js')
    .sass('client_src/app.scss', 'static/app.css')
    .setPublicPath('static');