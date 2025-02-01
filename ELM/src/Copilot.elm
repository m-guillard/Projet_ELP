module Copilot exposing (..)

import Parser exposing (..)
import Parser.Advanced exposing (..)

type Command
    = Forward Int
    | LeftTurn Int
    | RightTurn Int
    | Repeat Int (List Command)

commandParser : Parser Command
commandParser =
    oneOf
        [ forwardParser
        , leftParser
        , rightParser
        , repeatParser
        ]

forwardParser : Parser Command
forwardParser =
    succeed Forward
        |= (symbol "Forward" *> spaces *> int)

leftParser : Parser Command
leftParser =
    succeed LeftTurn
        |= (symbol "Left" *> spaces *> int)

rightParser : Parser Command
rightParser =
    succeed RightTurn
        |= (symbol "Right" *> spaces *> int)

repeatParser : Parser Command
repeatParser =
    succeed Repeat
        |= (symbol "Repeat" *> spaces *> int)
        |= (spaces *> symbol "[" *> spaces *> (commandParser `sepBy` (symbol "," *> spaces)) <* spaces <* symbol "]")

commandsParser : Parser (List Command)
commandsParser =
    succeed identity
        |= (symbol "[" *> spaces *> (commandParser `sepBy` (symbol "," *> spaces)) <* spaces <* symbol "]")

transformCommands : List Command -> List (Int, Int)
transformCommands commands =
    List.concatMap transformCommand commands

transformCommand : Command -> List (Int, Int)
transformCommand command =
    case command of
        Forward n ->
            [(n, 0)]

        LeftTurn n ->
            [(0, -n)]

        RightTurn n ->
            [(0, n)]

        Repeat times cmds ->
            List.concat (List.repeat times (transformCommands cmds))

main =
    let
        input = "[Repeat 2[Forward 30, Left 50], Repeat 2[Forward 20, Right 50]]"
        result = run commandsParser input
    in
    case result of
        Ok commands ->
            Debug.log "Transformed Commands" (transformCommands commands)

        Err err ->
            Debug.log "Parse Error" (Parser.Advanced.deadEndsToString err)