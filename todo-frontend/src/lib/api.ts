import axios from 'axios';

const API_BASE_URL = '/api';

export interface Todo {
  id: number;
  title: string;
  description: string;
  completed: boolean;
  CreatedAt?: string;
  UpdatedAt?: string;
}

export interface CreateTodoInput {
  title: string;
  description: string;
}

export interface UpdateTodoInput {
  title?: string;
  description?: string;
  completed?: boolean;
}

export const todoApi = {
  async getTodos(): Promise<Todo[]> {
    const response = await axios.get<Todo[]>(`${API_BASE_URL}/todos`);
    return response.data;
  },

  async getTodo(id: number): Promise<Todo> {
    const response = await axios.get<Todo>(`${API_BASE_URL}/todos/${id}`);
    return response.data;
  },

  async createTodo(input: CreateTodoInput): Promise<Todo> {
    const response = await axios.post<Todo>(`${API_BASE_URL}/todos`, input);
    return response.data;
  },

  async updateTodo(id: number, input: UpdateTodoInput): Promise<Todo> {
    const response = await axios.put<Todo>(`${API_BASE_URL}/todos/${id}`, input);
    return response.data;
  },

  async deleteTodo(id: number): Promise<void> {
    await axios.delete(`${API_BASE_URL}/todos/${id}`);
  },
};