package perldancer;
use Dancer ':syntax';

our $VERSION = '0.1';

get '/' => sub {
    template 'index';
};

get '/hello' => sub {
    "Hello There!! Nice Get!!";
};

get '/redirect' => sub {
    redirect '/';
};

true;
