import axios from 'axios';
// import { getCsrfToken } from './common';
const API_URL = 'http://localhost:8080/api/v1';

export const getEvent = async (eventId) => {
	return axios.get(`${API_URL}/events/${eventId}`);
}

export const createEvent = async (data) => {
	try {
		const response = await axios.post(`${API_URL}/events`, data, 
		{
			withCredentials: true 
		}
		);
		return {
			error: null,
			result: response.data
		};
	} catch (error) {
		if (error.response) {
			return {
				error: error.response.data,
				result: null
			};
		} else if (error.request) {
			return {
				error: 'No response received from server',
				result: null
			};
		} else {
			return {
				error: 'Error in setting up request',
				result: null
			};
		}
	}
}


export const getAllEvents = async (currentPage, perPage) => {
	const offset = (currentPage - 1) * perPage
	return axios.get(`${API_URL}/events/?limit=${perPage}&offset=${offset}`).then((response) => {
		return response.data
	});
}

export const getAllEventsByEventName = async (eventRule) => {
	return axios.get(`${API_URL}/events?rule=${eventRule}`).then((response) => {
		return response.data
	});
}
