import axios from 'axios';

const API_URL = 'http://localhost:8080/api/v1';

export const getEvent = (eventId) => {
  return axios.get(`${API_URL}/events/${eventId}`);
}

export const getAllEventsByEventName = (eventRule) => {
  return axios.get(`${API_URL}/events?rule=${eventRule}`);
}
