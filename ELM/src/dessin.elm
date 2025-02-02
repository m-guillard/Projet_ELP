module Dessin exposing (..)

import Svg exposing (Svg, svg, line)
import Svg.Attributes exposing (x1, y1, x2, y2, stroke, strokeWidth, viewBox, width, height)
import Parsing exposing (..)
import Debug exposing (log)

type alias Coord_deux_points =
    { x1 : Float
    , y1 : Float
    , x2 : Float
    , y2 : Float
    }

drawLine : Coord_deux_points -> Svg msg
drawLine { x1, y1, x2, y2 } =
    line
        [ Svg.Attributes.x1 (String.fromFloat x1)
        , Svg.Attributes.y1 (String.fromFloat y1)
        , Svg.Attributes.x2 (String.fromFloat x2)
        , Svg.Attributes.y2 (String.fromFloat y2)
        , stroke "blue"
        , strokeWidth "2"
        ]
        []


display : List Parsing.Point -> List (Svg msg)
display points =
    case points of
        [] -> []
        [ _ ] -> []
        p1 :: p2 :: rest -> 
            drawLine { x1 = p1.x, y1 = p1.y, x2 = p2.x, y2 = p2.y } :: display (p2 :: rest)

transformLinesToPoints : List Coord_deux_points -> List Parsing.Point
transformLinesToPoints lines =
    List.concatMap 
        (\{ x1, y1, x2, y2 } -> 
            [ { x = x1, y = y1 }, { x = x2, y = y2 } ] 
        )
        lines
