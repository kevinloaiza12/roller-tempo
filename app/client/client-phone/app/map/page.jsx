'use client';

import {Helmet} from "react-helmet";
import placeholder from './map.jpg'
import './map.css'
import Image from 'next/image'
import 'leaflet/dist/leaflet.css';

export const revalidate = 5;

const atracciones2 = [
  {
    "capacity": "50",
    "description": "Fire",
    "duration": "60",
    "name": "a",
    "nextTurn": 50,
  },
  {
    "capacity": "20",
    "description": "water",
    "duration": "6",
    "name": "b ",
    "nextTurn": 501,

  },
  {
    "capacity": "1",
    "description": "Earth",
    "duration": "1",
    "name": "c ",
    "nextTurn": 10,
  }
]

export default function Atraccion() {

    return (
        <div>
            <Image id="image" src={placeholder}/>
        </div>
    )
  }
  