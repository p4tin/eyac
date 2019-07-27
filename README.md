# EYAC (Promunced like `Yak`)
#### Environment Yaml Augmented Configuration

                            _,,,_
                        .-'`  (  '.
                     .-'    ,_  ;  \___      _,
                 __.'    )   \'.__.'(:;'.__.'/
         __..--""       (     '.__{':');}__.'
       .'         (    ;    (   .-|` '  |-.
      /    (       )     )      '-p     q-'
     (    ;     ;          ;    ; |.---.|
     ) (              (      ;    \ o  o)
     |  )     ;       |    )    ) /'.__/
     )    ;  )    ;   | ;       //
     ( )             _,\    ;  //
     ; ( ,_,,-~""~`""   \ (   //
      \_.'\\_            '.  /<_
       \\_)--\             \ \--\
       )--\""`             )--\"`
       `""`                `""`

Library that laods a yaml file as your configuration and then augments it with environment variables if supplied or uses the defaults.

example:

config.yaml

HOST: <%= ENV['RMQ_HOST'] %>, localhost
PORT: <%= ENV['RMQ_PORT'] %>, 8080

When you call LoadConfig if the environment variables `RMQ_HOST` or `RMQ_PORT` are set they will override the defaults `localhost` and `8080`

See the examples directory for a more complex example.
