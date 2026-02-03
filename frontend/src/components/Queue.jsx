import React, { useEffect, useState } from 'react'
import { fetchQueue, addToQueue } from '../api'

export default function Queue() {
  const [queue, setQueue] = useState([])
  const [name, setName] = useState('')
  const [song, setSong] = useState('')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)

  async function load() {
    setLoading(true)
    setError(null)
    try {
      const data = await fetchQueue()
      setQueue(data)
    } catch (e) {
      setError(e.message)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => { load() }, [])

  async function handleAdd(e) {
    e.preventDefault()
    setError(null)
    try {
      await addToQueue({ name, song })
      setName('')
      setSong('')
      await load()
    } catch (e) {
      setError(e.message)
    }
  }

  return (
    <div className="queue">
      <form onSubmit={handleAdd} className="add-form">
        <input placeholder="Nombre" value={name} onChange={e => setName(e.target.value)} required />
        <input placeholder="Canción" value={song} onChange={e => setSong(e.target.value)} required />
        <button type="submit">Agregar</button>
      </form>

      {loading ? <p>Cargando...</p> : (
        <ul className="list">
          {queue.length === 0 ? <li>No hay entradas</li> : queue.map((q, i) => (
            <li key={i}><strong>{q.name}</strong> — {q.song}</li>
          ))}
        </ul>
      )}
      {error && <p className="error">{error}</p>}
    </div>
  )
}