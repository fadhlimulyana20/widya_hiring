import { Auth } from "@/models/user";
import { PayloadAction, createSlice } from "@reduxjs/toolkit";

export interface AuthState {
    auth: Auth
}

const initialState: AuthState = {
    auth: {
        token: {
            access: '',
            refresh: '',
            timeout: ''
        },
        user: {
            created_at: '',
            email: '',
            name: '',
            roles: []
        }
    }
}

export const authSlice = createSlice({
    name: 'auth',
    initialState,
    reducers: {
        setAuth: (state, action: PayloadAction<Auth>) => {
            state.auth = action.payload
        }
    }
})

export const {setAuth} = authSlice.actions
export default authSlice.reducer
