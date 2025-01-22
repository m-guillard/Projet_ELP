module Parsing exposing (..)

import Parser exposing (suceed, token, int, spaces, lazy, run)

type perso
read : String -> perso