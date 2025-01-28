module Parsing exposing (..)

import Parser exposing (succeed, token, int, spaces, lazy, run)

type alias Instructions =
    { nb : Int
    , instr : List (List Int)
    }

type Direction =
    | Left
    | Right

type alias Expression =
    ( String, Int )

intParser : ParserInt

read : String -> Instructions
read

calcul : Int -> Int -> Dict String Int

parser : Parser.Parser (Instructions)
parser =
    lazy(\_ ->
        spaces
            |> Parser.andThen(\_ -> token "[")
            |> Parser.andThen(\_ -> token "Repeat")
                spaces |> Parser.andThen (\_ ->
                    int |> Parser.andThen (\value ->
                        succeed(value)
                    )
                )
            |> Parser.andThen (\_ -> token Direction)
                spaces |> Parser.andThen (\_ ->
                        int |> Parser.andThen (\value ->
                            succeed(value)
                        )
                    )
            |> Parser.andThen (\_ -> parser)
            |> Parser.andThen(\_ -> token("]"))
    
    
    )
    