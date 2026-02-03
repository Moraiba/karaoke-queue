const API_BASE = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'

export async function fetchQueue() {
  const res = await fetch(`${API_BASE}/queue`)
  if (!res.ok) throw new Error('Error fetching queue')
  return res.json()
}

export async function addToQueue(item) {
  const res = await fetch(`${API_BASE}/queue`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(item)
  })
  if (!res.ok) throw new Error('Error adding to queue')
  return res.json()
}