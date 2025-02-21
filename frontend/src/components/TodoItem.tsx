export interface Todo {
    id: number;
    task: string;
    status: 'pending' | 'completed';
    date: string;
}

interface TodoItemProps {
    todo: Todo;
}

const TodoItem: React.FC<TodoItemProps> = ({ todo }) => {
    return (
        <div className="border p-4 rounded">
            {todo.task} - {todo.status} - {todo.date}
        </div>
    );
};

export default TodoItem;