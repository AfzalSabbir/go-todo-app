import {createBrowserRouter} from "react-router-dom";

import Home from "../components/Home.tsx";
import TodoList from "../components/TodoList.tsx";
import TodoCreate from "../components/TodoCreate.tsx";
import TodoEdit from "../components/TodoEdit.tsx";
import TodoDelete from "../components/TodoDelete.tsx";

const routes = createBrowserRouter([
    {path: '/', element: <Home/>},
    {path: '/about', element: <Home/>},
    {path: '/todos', element: <TodoList/>}, // List all todos
    {path: '/todos/create', element: <TodoCreate/>}, // Create a new todo
    {path: '/todos/edit/:id', element: <TodoEdit/>}, // Edit a specific todo by ID
    {path: '/todos/delete/:id', element: <TodoDelete/>}, // Delete a specific todo by ID
    {path: '*', element: <Home/>} // Catch-all for unknown routes
]);

export default routes;
