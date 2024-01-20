// service.js
import apiClient from '../api';

// GET 요청
const getTodo = async () => {
  try {
    const response = await apiClient.get('/todos');
    return response.data;
  } catch (error) {
    console.error('GET Request Error:', error);
    throw error;
  }
};
// GET 요청
const SearchTodo = async (keyword) => {
  try {
    const response = await apiClient.get('/todos?id='+keyword);
    return response.data;
  } catch (error) {
    console.error('GET Request Error:', error);
    throw error;
  }
};
// POST 요청
const postTodo = async (title) => {
  const todo = {title: title};
  try {
    const response = await apiClient.post('/todos', todo);
    return response.data;
  } catch (error) {
    console.error('POST Request Error:', error);
    throw error;
  }
};

const putTodo = async (id, completed) => {
  const data = {
    ID: id,
    Completed: completed
  };

  try {
    const response = await apiClient.put('/todos', data);
    return response.data;
  } catch (error) {
    console.error('PUT Request Error:', error);
    throw error;
  }
};

const deleteTodo = async (id) => {
  try {
    const response = await apiClient.delete('/todos', { data: { id } });
    return response.data;
  } catch (error) {
    console.error('DELETE Request Error:', error);
    throw error;
  }
};


export { getTodo, postTodo, putTodo, deleteTodo };
