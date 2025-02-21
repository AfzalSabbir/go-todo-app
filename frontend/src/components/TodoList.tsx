import React, {useState} from 'react';

import TodoItem, {Todo} from "./TodoItem.tsx";

const TodoList: React.FC = () => {
    const [todos] = useState<Todo[]>([
        {id: 1, task: 'Buy groceries', status: 'pending', date: '2025-02-01'},
        {id: 2, task: 'Complete project', status: 'completed', date: '2025-02-05'},
        {id: 3, task: 'Attend meeting', status: 'pending', date: '2025-02-10'},
    ]);

    const [search, setSearch] = useState<string>('');
    const [pending, setPending] = useState<string>('');
    const [startDate, setStartDate] = useState<string>('');
    const [endDate, setEndDate] = useState<string>('');

    const handleSearchChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setSearch(e.target.value);
    };

    const handlePendingChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        setPending(e.target.value);
    };

    const handleStartDateChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setStartDate(e.target.value);
    };

    const handleEndDateChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setEndDate(e.target.value);
    };

    const filteredTodos = todos.filter(todo => {
        const taskMatches = todo.task.toLowerCase().includes(search.toLowerCase());
        const statusMatches = !pending || todo.status === pending;
        const dateMatches =
            (!startDate || new Date(todo.date) >= new Date(startDate)) &&
            (!endDate || new Date(todo.date) <= new Date(endDate));

        return taskMatches && statusMatches && dateMatches;
    });

    return (
        <div className="container mx-auto p-6">
            {/* Filters Section */}
            <div className="flex space-x-4 mb-6">
                {/* Search */}
                <div className="flex items-center">
                    <label htmlFor="search" className="mr-2">Search</label>
                    <input
                        type="text"
                        id="search"
                        className="border p-2 rounded"
                        placeholder="Search tasks..."
                        value={search}
                        onChange={handleSearchChange}
                    />
                </div>

                {/* Select Pending Tasks */}
                <div className="flex items-center">
                    <label htmlFor="pending" className="mr-2">Pending</label>
                    <select
                        id="pending"
                        className="border p-2 rounded"
                        value={pending}
                        onChange={handlePendingChange}
                    >
                        <option value="">All</option>
                        <option value="pending">Pending</option>
                        <option value="completed">Completed</option>
                    </select>
                </div>

                {/* Date Range Filter */}
                <div className="flex items-center">
                    <label htmlFor="start-date" className="mr-2">Start Date</label>
                    <input
                        type="date"
                        id="start-date"
                        className="border p-2 rounded"
                        value={startDate}
                        onChange={handleStartDateChange}
                    />
                    <label htmlFor="end-date" className="ml-4 mr-2">End Date</label>
                    <input
                        type="date"
                        id="end-date"
                        className="border p-2 rounded"
                        value={endDate}
                        onChange={handleEndDateChange}
                    />
                </div>
            </div>

            {/* Todo List Section */}
            <div id="todo-list" className="space-y-4">
                {filteredTodos.map(todo => (
                    <TodoItem key={todo.id} todo={todo}/>
                ))}
            </div>
        </div>
    );
};

export default TodoList;
