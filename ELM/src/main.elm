module Main exposing (..)

import Browser
import Html exposing (..)
import Html.Events exposing (..)
import Html.Attributes exposing(style, placeholder)
import Svg exposing(svg, line)
import Svg.Attributes exposing(viewBox, width, height)
-- import Parsing exposing (..)
-- import Dessin exposing (..)

-- MAIN

main =
    Browser.sandbox {init = init, update = update, view = view}

-- MODEL

type alias Model = Int

init : Model
init =
    0

-- UPDATE

type Msg
    = Draw

update : Msg -> Model -> Model
update msg model = 
    case msg of
        Draw ->
            model + 1

-- VIEW

stylePage : List(Html.Attribute Msg)
stylePage =
    [ style "box-sizing" "border-box"
    , style "width" "100%"
    , style "max-width" "500px"
    , style "margin" "auto"
    , style "padding" "1rem"
    , style "display" "grid"
    , style "grid-gap" "30px"
    ]

view : Model -> Html Msg
view model =
    div [ style "font-family" "-apple-system, BlinkMacSystemFont, Segoe UI, Roboto, Open Sans, Ubuntu, Fira Sans, Helvetica Neue, sans-serif"
        , style "margin" "0"
        , style "min-height" "100vh"
        ]
        [ div
            stylePage
            [ text "Type in your code below:"]
        , div
            stylePage
            [ input
                [ placeholder "example: [Repeat 360 [Forward 1, Left 1]]"] 
                []
            , button
                [ onClick Draw ]
                [ text "Draw" ]
            , svg
                [ style "width" "100%"
                , style "max-width" "500px"
                , style "border" "1px solid #D1C4E9"
                , viewBox "0 0 500 500"
                , width "500"
                , height "500"
                ]
                []
                --[ display model ]
            ]
        ]