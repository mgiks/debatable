import type { SyntheticEvent } from 'react'
import './CreatePostPage.scss'

function CreatePostPage() {
  return (
    <div className='page'>
      <Form />
    </div>
  )
}

function Form() {
  const handleSubmit = async (ev: SyntheticEvent) => {
    ev.preventDefault()

    const formData = new FormData(ev.target as HTMLFormElement)
    const data = Object.fromEntries(formData)

    try {
      const response = await fetch('http://localhost:8080/v1/posts/', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      })

      const result = await response.json()
      console.log(result)
    } catch (err) {
      console.log(err)
    }
  }

  return (
    <form
      className='create-task-form'
      action={'http://localhost:8080/v1/posts/'}
      method='POST'
      name='createTaskForm'
      onSubmit={handleSubmit}
    >
      <label htmlFor='title'>Title</label>
      <input type='text' name='title' id='title' />
      <label htmlFor='body'>Body</label>
      <input type='text' name='body' id='body' />
      <button type='submit'>Create post</button>
    </form>
  )
}

export default CreatePostPage
