import './SideBar.scss'
import { Link } from 'react-router'

function SideBar() {
  return (
    <nav className='sidebar'>
      <ul>
        <li>
          <Link to={'/'}>
            Home
          </Link>
        </li>
        <li>
          <Link to={'/posts/new'}>
            Create Post
          </Link>
        </li>
      </ul>
    </nav>
  )
}

export default SideBar
