module Dessin exposing (..)

import Svg exposing (Svg, svg, line)
import Svg.Attributes exposing (x1, y1, x2, y2, stroke, strokeWidth, viewBox, width, height)

-- Définir un alias pour un point (les 4 coordonnées d'une ligne)
type alias Point =
    { x1 : Float
    , y1 : Float
    , x2 : Float
    , y2 : Float
    }

-- Fonction pour afficher une liste de lignes sous forme d'éléments SVG
display : List Point -> List (Svg msg)
display lines =
    List.map drawLine lines

-- Fonction pour dessiner une ligne individuelle
drawLine : Point -> Svg msg
drawLine point =
    line
        [ x1 (String.fromFloat point.x1)
        , y1 (String.fromFloat point.y1)
        , x2 (String.fromFloat point.x2)
        , y2 (String.fromFloat point.y2)
        , stroke "black" -- Couleur de la ligne
        , strokeWidth "2" -- Épaisseur de la ligne
        ]
        []

-- Exemple de test pour générer un canevas SVG avec des lignes
testDisplay : Svg msg
testDisplay =
    let
        -- Liste de lignes à afficher avec des records
        lines =
            [ { x1 = 50, y1 = 50, x2 = 150, y2 = 50 }
            , { x1 = 150, y1 = 50, x2 = 150, y2 = 150 }
            , { x1 = 150, y1 = 150, x2 = 50, y2 = 150 }
            , { x1 = 50, y1 = 150, x2 = 50, y2 = 50 }
            , { x1 = 50, y1 = 50, x2= 150, y2 = 150}
            , { x1 = 150, y1 = 150, x2= 200, y2 = 200}
            ]
    in
    svg
        [ viewBox "0 0 200 200"
        , width "200"
        , height "200"
        ]
        (display lines)

-- Valeur main pour rendre l'élément SVG
main =
    testDisplay
