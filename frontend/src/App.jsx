import React from 'react'
import Queue from './components/Queue'

export default function App() {
  return (
    <div className="app">
      <header>
        <h1>Karaoke Queue</h1>
      </header>
      <main>
        <Queue />
      </main>
    </div>
  )
}