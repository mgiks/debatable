import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.scss'
import App from './App.tsx'
import { createBrowserRouter, RouterProvider } from 'react-router'
import CreatePostPage from './features/posts/components/CreatePostPage.tsx'

const router = createBrowserRouter([
  {
    path: '/',
    Component: App,
  },
  {
    path: '/posts/new',
    Component: CreatePostPage,
  },
])

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
)
