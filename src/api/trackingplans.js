import axios from "axios";

const API_URL = 'http://localhost:8080/api/v1';


export const createTrackingPlan = async (data) => {
	try {
		const response = await axios.post(`${API_URL}/tracking-plans`, data,
		{
			withCredentials: true 
		});
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

export const getAllTrackingPlans = (currentPage, perPage) => {
	const offset = (currentPage-1)*perPage
	return axios.get(`${API_URL}/tracking-plans/?limit=${perPage}&offset=${offset}`).then((response) => {
		return response.data
	});
}

export const getTrackingPlan = (id) => {
	return axios.get(`${API_URL}/tracking-plans/${id}`).then((response) => {
		return response.data
	});
}

export const updateTrackingPlan = async (id, data) => {
	try {
		const response = await axios.put(`${API_URL}/tracking-plans/${id}`, data,
		{
			withCredentials: true 
		});
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
