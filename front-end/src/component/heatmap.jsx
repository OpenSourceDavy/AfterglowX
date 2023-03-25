/* global google */
import React from 'react'
import GoogleMapReact from 'google-map-react'
import {useState} from "react";
import {useEffect} from "react";

export default function HeatMapshow(props){
    const [cordinates , setCordinates] = useState([43.653225, -79.383186]);

    const defaultProps = {

        zoom: 10
    };
    const heatMapData={
        positions: [
            {lat: props.userLat, lng: props.userLng},

        ],
        options: {
            radius: 15,
            dissipating:true,
            opacity:.5,

            // scaleControl:false
        }}
    useEffect(() => {
        setCordinates([props.userLat
    , props.userLng])}, [])

    return (
        // Important! Always set the container height explicitly
        <div style={{ height: '100vh', width: '100%' }}>
            <GoogleMapReact
                bootstrapURLKeys={{
                    key: '',
                    libraries: ['visualization'],
                }}
                defaultCenter={cordinates}
                defaultZoom={defaultProps.zoom}
                options={{ fullscreenControl: false, gestureHandling: 'cooperative', scrollwheel: false}}
                // heatmapLibrary={true}
                heatmap={heatMapData}
            >
            </GoogleMapReact>
        </div>
    );
}

