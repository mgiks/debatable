import { Link } from 'react-router'
import './App.scss'

function App() {
  return (
    <div>
      <Link to={'/posts/new'}>
        <button>
          Create Post
        </button>
      </Link>
    </div>
  )
}

export default App
