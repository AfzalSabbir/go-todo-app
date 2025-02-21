export interface Todo {
    id: number;
    task: string;
    status: 'pending' | 'completed';
    created_at: string;
}

interface TodoItemProps {
    todo: Todo;
}

const TodoItem: React.FC<TodoItemProps> = ({ todo }) => {
    return (
        <div className="border p-4 rounded">
            {todo.task} - {todo.status} - {todo.created_at}
        </div>
    );
};

export default TodoItem;