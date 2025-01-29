module Parsing exposing (..)

import Parser exposing (succeed, token, int, spaces, lazy, run)
import Array exposing (empty)

type alias Instructions =
    { nb : Int
    , instr : List (List Int)
    }

type InstrIntermediare a
    = Chaine String
    | Instruction { nb : Int, instr : InstrIntermediare a}

findParser : Parser.Parser String
findParser =
    Parser.oneOf
        [ token "["
        , token "]"
        ]

fonction1 : InstrIntermediare String -> String -> (InstrIntermediare String, String)
fonction1 instr inputString =
    case run findParser inputString of
        Ok _ -> fonction1 (transformationCrochet Instruction)
        Err _ ->
            Instruction { nb = 1, instr = instr}
            
        
transformationCrochet : InstrIntermediare -> InstrIntermediare

type Direction =
    | Left
    | Right

type alias Expression =
    ( String, Int )

intParser : Parser.Parser Int
intParser =
    int

directionParser : Parser.Parser String
directionParser =
    Parser.oneOf
        [ token "Left"
        , token "Right"
        , token "Repeat"
        ]

repeatParser : Parser.Parser String
repeatParser =
    Parser.oneOf
        token "Repeat"

expressionParser : Parser.Parser Expression*
expressionParser =
    |= spaces

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
    