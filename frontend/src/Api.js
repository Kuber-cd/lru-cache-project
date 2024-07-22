import axios from 'axios';

const API_URL = 'http://localhost:8080';

export const setCache = async (key, value, expiration) => {
  try {
    await axios.get(`${API_URL}/set`, {
      params: { key, value, expiration }
    });
  } catch (error) {
    console.error("Error setting cache:", error);
  }
};

export const getCache = async (key) => {
  try {
    const response = await axios.get(`${API_URL}/get`, { params: { key } });
    return response.data;
  } catch (error) {
    console.error("Error getting cache:", error);
    return null;
  }
};

export const deleteCache = async (key) => {
  try {
    await axios.get(`${API_URL}/delete`, { params: { key } });
  } catch (error) {
    console.error("Error deleting cache:", error);
  }
};