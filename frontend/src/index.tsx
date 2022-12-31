import Footer from 'components/layouts/Footer'
import Header from 'components/layouts/Header'

import Login from 'pages/Login'
import SignUp from 'pages/SignUp'

import Task from 'pages/task/TaskId'
import TaskEdit from 'pages/task/TaskIdEdit'
import TaskNew from 'pages/task/TaskNew'
import TaskList from 'pages/tasks/TaskList'

import Top from 'pages/Top'

import User from 'pages/user/UserId'
import UserEdit from 'pages/user/UserIdEdit'
import UserNew from 'pages/user/UserNew'

import ReactDOM from 'react-dom/client'
import { BrowserRouter, Route, Routes } from 'react-router-dom'

import 'styles/index.scss'
import reportWebVitals from './reportWebVitals'

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
)

root.render(
  <BrowserRouter>
    <Header />

    <Routes>
      <Route path="/" element={<Top />} />

      {/* task */}
      <Route path="/tasks" element={<TaskList />} />
      <Route path="/task/:id" element={<Task />} />
      <Route path="/task/:id/edit" element={<TaskEdit />} />
      <Route path="/task/new" element={<TaskNew />} />

      {/* auth */}
      <Route path="/signup" element={<SignUp />} />
      <Route path="/login" element={<Login />} />

      {/* user */}
      <Route path="/user/:id" element={<User />} />
      <Route path="/user/:id/edit" element={<UserEdit />} />
      <Route path="/user/new" element={<UserNew />} />

      {/* 404 */}
      <Route path="*" element={<div>404</div>} />
    </Routes>

    <Footer />
  </BrowserRouter>
)

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals()
