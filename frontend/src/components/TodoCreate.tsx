import {useEffect, useState} from "react";
import axios from "axios";
import {useParams} from "react-router-dom";
import {Button} from "./ui/button.tsx";
import {Input} from "./ui/input.tsx";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "./ui/select.tsx";

export default function TodoCreate() {
    const params = useParams();

    const [task, setTask] = useState({
        task: '',
        status: 'pending'
    });

    useEffect(() => {
        axios
            .get(`http://localhost:8080/todos/${params.id}`)
            .then((r) => {
                console.log("response", r.data);
                setTask(r.data.todo)
            });
    }, [])

    const addTask = () => {
        axios
            .post("http://localhost:8080/todos", {
                ...task
            })
            .then((r) => {
                setTask(r.data.todo)
            });
    };

    const updateTask = () => {
        axios
            .put(`http://localhost:8080/todos/${params.id}`, {
                task: task.task,
                status: task.status
            })
            .then((r) => {
                setTask(r.data.todo)
            });
    };

    return (
        <div className="container mx-auto p-6 max-w-md">
            <h1 className="text-2xl font-bold mb-4">To-Do List</h1>

            <div className="flex gap-2">
                <Input
                    type="text"
                    className="w-full"
                    placeholder="Add a task..."
                    value={task.task}
                    onChange={(e) => setTask({...task, task: e.target.value})}
                />

                <Select
                    value={task.status}
                    onValueChange={(value) => {
                        console.log('value', value);
                        setTask({...task, status: value})
                    }}
                >
                    <SelectTrigger className="border p-2 rounded">
                        <SelectValue placeholder="Select status"/>
                    </SelectTrigger>
                    <SelectContent
                        className="border p-2 rounded">
                        <SelectItem value="pending">Pending</SelectItem>
                        <SelectItem value="complete">Complete</SelectItem>
                    </SelectContent>
                </Select>

                <Button
                    className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
                    onClick={!params.id ? addTask : updateTask}
                >
                    {!params.id ? 'Add' : 'Update'}
                </Button>
            </div>
        </div>
    );
}
