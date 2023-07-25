import { createStore, Store } from 'vuex'
import { AxiosError } from 'axios';

import { Cart } from '@/typings';

import ApiService from '@/services/api.service';
import ProductModel from '@/models/product.model';

const cart: Cart = JSON.parse(localStorage.getItem('cart') || '{}');

function getIdString<T extends {id: number}>(obj: T): string {
  return '_' + obj.id.toString();
}

type State = {
  cart: Cart;
  auth: {
    checked: boolean,
    isLoggedIn: boolean,
    user?: {
      id: number,
      first_name: string,
      last_name: string,
      email: string,
    }
  }
}

const store: Store<State> = createStore({
  state: {
    cart,
    auth: {
      checked: false,
      isLoggedIn: false, 
    }
  } as State,
  getters: {
    cartAsArray: state => {
      return Object.values(state.cart);
    },
    numberOfItemsInCart: state => {
      return Object.values(state.cart).reduce((acc, item) => {
        return acc + item.quantity;
      }, 0);
    },
    cartTotal: state => {
      return Object.values(state.cart).reduce((acc, item) => {
        return acc + item.quantity * item.product.price;
      }, 0);
    },
    userFullname: state => {
      if(!state.auth.user) return '';
      return state.auth.user.first_name + ' ' + state.auth.user.last_name;
    }
  },
  mutations: {
    addToCart(state, product: ProductModel) {
      
      if(state.cart[getIdString(product)]) {
        state.cart[getIdString(product)].quantity++;
      } else {
        state.cart[getIdString(product)] = {
          product,
          quantity: 1,
        };
      }

    },
    removeFromCart(state, product: ProductModel) {
      delete state.cart[getIdString(product)];
    },
    setAuthChecked(state, value: boolean) {
      state.auth.checked = value;
    },
    setAuthLoggedIn(state, value: boolean) {
      state.auth.isLoggedIn = value;
    },
    setAuthUser(state, user: any) {
      state.auth.user = user;
    }
  },
  actions: {
    async checkAuth({ commit }) {
      
      const user = await ApiService.get('/auth/whoami/').catch((error: AxiosError) => {
        if(error.response && error.response.status === 403) return null
      });

      commit('setAuthChecked', true);
      commit('setAuthLoggedIn', !!user);
      commit('setAuthUser', user);

    },
    login({ commit }, credentials: {email: string, password: string}) {
      return ApiService.post('/auth/login/', credentials).then(user => {
        commit('setAuthLoggedIn', true);
        commit('setAuthUser', user);
        return user;
      });
    },
    logout({ commit }) {
      return ApiService.post('/auth/logout/').then(() => {
        commit('setAuthLoggedIn', false);
        commit('setAuthUser', null);
        ApiService.csrf_token = '';
      });
    },
    changePassword({ commit }, credentials: {old_password: string, new_password: string}) {
      return ApiService.post('/auth/change-password/', credentials).then((data: any) => { 
        commit('setAuthLoggedIn', false);
        commit('setAuthUser', null);
        ApiService.csrf_token = '';
        return data;
      });
    },
  },
  modules: {
  }
});

store.dispatch('checkAuth');

store.subscribe((mutation, state) => {
  localStorage.setItem('cart', JSON.stringify(state.cart));
});

export default store;
