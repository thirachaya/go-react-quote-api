import { useState } from "react"

type Quote = {
  id: number
  text: string
  author: string
}

function App() {

  const [quote, setQuote] = useState<Quote | null>(null)

  const fetchQuote = async () => {
    const res = await fetch(`${import.meta.env.VITE_API_URL}/quotes/random`)
    const data = await res.json()
    setQuote(data)
  }

  return (
    <div style={{ padding: 40 }}>
      <h1>Quote of the Day</h1>

      <button onClick={fetchQuote}>
        Get Random Quote
      </button>

      {quote && (
        <div style={{ marginTop: 20 }}>
          <p>"{quote.text}"</p>
          <p>- {quote.author}</p>
        </div>
      )}

    </div>
  )
}

export default App