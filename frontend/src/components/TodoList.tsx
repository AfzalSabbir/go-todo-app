import React, { useEffect, useState } from 'react';
import axios from "axios";
import {NavLink} from "react-router-dom";

import { Todo } from "./TodoItem.tsx";

const TodoList: React.FC = () => {
    const [todos, setTodos] = useState<Todo[]>([]);
    const [search, setSearch] = useState<string>('');
    const [pending, setPending] = useState<string>('');
    const [startDate, setStartDate] = useState<string>('');
    const [endDate, setEndDate] = useState<string>('');

    useEffect(() => {
        axios.get('http://localhost:8080/todos').then((r) => {
            console.log('response', r.data.todos);
            setTodos(r.data.todos);
        });
    }, []);

    const handleSearchChange = (e: React.ChangeEvent<HTMLInputElement>) => setSearch(e.target.value);
    const handlePendingChange = (e: React.ChangeEvent<HTMLSelectElement>) => setPending(e.target.value);
    const handleStartDateChange = (e: React.ChangeEvent<HTMLInputElement>) => setStartDate(e.target.value);
    const handleEndDateChange = (e: React.ChangeEvent<HTMLInputElement>) => setEndDate(e.target.value);

    const handleEdit = (id: number) => {
        console.log("Edit", id);
    };

    const handleDelete = (id: number) => {
        console.log("Delete", id);
    };

    const handleView = (id: number) => {
        console.log("View", id);
    };

    const filteredTodos = todos.filter(todo => {
        const taskMatches = todo.task.toLowerCase().includes(search.toLowerCase());
        const statusMatches = !pending || todo.status === pending;
        const dateMatches =
            (!startDate || new Date(todo.created_at) >= new Date(startDate)) &&
            (!endDate || new Date(todo.created_at) <= new Date(endDate));
        return taskMatches && statusMatches && dateMatches;
    });

    return (
        <div className="container mx-auto p-6">
            {/* Filters Section */}
            <div className="flex space-x-4 mb-6">
                <input type="text" className="border p-2 rounded" placeholder="Search tasks..." value={search} onChange={handleSearchChange} />
                <select className="border p-2 rounded" value={pending} onChange={handlePendingChange}>
                    <option value="">All</option>
                    <option value="pending">Pending</option>
                    <option value="completed">Completed</option>
                </select>
                <input type="date" className="border p-2 rounded" value={startDate} onChange={handleStartDateChange} />
                <input type="date" className="border p-2 rounded" value={endDate} onChange={handleEndDateChange} />
                <NavLink to="/todos/create" className="ml-auto bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600">
                    Create
                </NavLink>
            </div>

            {/* Todo List Table */}
            <table className="table-auto w-full border-collapse border border-gray-300">
                <thead>
                <tr className="bg-gray-200">
                    <th className="border border-gray-300 p-2">Task</th>
                    <th className="border border-gray-300 p-2">Status</th>
                    <th className="border border-gray-300 p-2">Created At</th>
                    <th className="border border-gray-300 p-2">Actions</th>
                </tr>
                </thead>
                <tbody>
                {filteredTodos.map(todo => (
                    <tr key={todo.id} className="border border-gray-300">
                        <td className="border border-gray-300 p-2">{todo.task}</td>
                        <td className="border border-gray-300 p-2">{todo.status}</td>
                        <td className="border border-gray-300 p-2">{new Date(todo.created_at).toLocaleDateString()}</td>
                        <td className="border border-gray-300 p-2 flex space-x-2">
                            <button onClick={() => handleEdit(todo.id)} className="bg-blue-500 text-white px-2 py-1 rounded">Edit</button>
                            <button onClick={() => handleDelete(todo.id)} className="bg-red-500 text-white px-2 py-1 rounded">Delete</button>
                            <button onClick={() => handleView(todo.id)} className="bg-green-500 text-white px-2 py-1 rounded">View</button>
                        </td>
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    );
};

export default TodoList;