import Image from 'next/image'
import { Inter } from 'next/font/google'
import styles from './page.module.css'
import fetch from "node-fetch";

const inter = Inter({ subsets: ['latin'] })

const getPlanets = async () => {
  const res = await fetch('https://swapi.dev/api/planets');
  
  
  return res.json();
}

const getAtracction = () => {
  fetch("http://127.0.0.1:3000/api/attractioninfo/Divertido")
  .then( (response) => response.json())
  .then((data)=> console.log(data))// output will be the required data
  .catch( (error) => console.log(error))
}

async function getData() {
  let response = await
  fetch("http://127.0.0.1:3000/api/attractioninfo/Divertido")
  let data = await response.json()
  return data;
  }

export default async function Home() {
  //const planets = await getPlanets();
   const data = await getData();
  return (
    <main className={styles.main}>
      <h1>Hola</h1>
      {/* <ul>
        {planets.results.map((planer, index) => (
          <li key={index}>{planer.name}</li>
        ))}
      </ul> */}
      <p>{data.name}</p>
    </main>
  )
}
