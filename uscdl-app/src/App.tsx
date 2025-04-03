import { useState } from 'react'

import USCDLEditor from './USCDLEditor'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <USCDLEditor />
    </>
  )
}

export default App
