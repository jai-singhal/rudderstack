import axios from "axios";

const API_URL = 'http://localhost:8080/api/v1';

export const createTrackingPlan = async (data) => {
  try {
    const response = await axios.post(`${API_URL}/tracking-plans/`, data, {
      headers: {
        'Content-Type': 'application/json'
      }
    });
    return response.data;
  } catch (error) {
    throw new Error(error.message);
  }
}

export const getAllTrackingPlans = () => {
  return axios.get(`${API_URL}/tracking-plans/`).then((response)=>{
    return response.data
  });
}

export const getTrackingPlan = (id) => {
  return axios.get(`${API_URL}/tracking-plans/${id}`).then((response)=>{
    return response.data
  });
}
