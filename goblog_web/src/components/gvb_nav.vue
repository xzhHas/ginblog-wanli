<template>
  <div :class="{
    gvb_nav: true,
    show: data.is_show,
  }">
    <div class="gvb_nav_container">
      <div class="logo">
        <div>WANLI</div>
        <div>Always Lucky</div>
      </div>
      <div class="left">
        <span v-for="item in store.navList" :key="item.id">
          <template v-if="item.path.startsWith('http')">
            <a :href="item.path" target="_blank">{{ item.title }}</a>
          </template>
          <template v-else>
            <router-link :to="item.path">{{ item.title }}</router-link>
          </template>
        </span>

        <span class="search">
          <GVBTextSearch />
        </span>
      </div>
      <div class="right">
        <GVBTheme class="gvb_theme"></GVBTheme>
        <span class="login_btn" v-if="store.userInfo.role === 0"><router-link
            :to="{ name: 'login', query: { redirect_url: $route.path } }">登录</router-link></span>
        <GVBUserInfo v-if="store.userInfo.role !== 0"></GVBUserInfo>
      </div>
    </div>
  </div>
</template>

<script setup>
import GVBTheme from "@/components/gvb_theme.vue";
import GVBUserInfo from "@/components/gvb_user_info.vue";
import { useStore } from "@/stores/store";
import { reactive } from "vue";
import GVBTextSearch from "@/components/gvb_text_search.vue";


const store = useStore()
const props = defineProps({
  is_show: {
    type: Boolean,
    default: false,
  }
})

const data = reactive({
  is_show: false,
})

// 顶栏是否一直显示
async function getInit() {
  if (props.is_show) {
    data.is_show = true
  } else {
    window.addEventListener("scroll", scroll)
  }
}

getInit()

// 页面滚动监控
function scroll() {
  let top = document.documentElement.scrollTop
  if (top >= 400) {
    data.is_show = true
  } else {
    data.is_show = false
  }
}
</script>

<style lang="scss">
.gvb_nav {
  position: fixed;
  top: 0;
  width: 100%;
  //background-color: transparent;
  height: 60px;
  //box-shadow: 1px 1px 5px #0003;
  display: flex;
  justify-content: center;
  font-size: 16px;
  z-index: 10;
  color: white;

  .gvb_nav_container {
    width: 1200px;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .logo {
    width: 10%;
    transform: scale(0.7);

    div:first-child {
      font-size: 26px;
    }

    div:last-child {
      font-size: 12px;
    }
  }

  .left {
    width: 65%;
    display: flex;

    span {
      margin-right: 30px;
    }

    .search i {
      cursor: pointer;
    }
  }

  a {
    color: white;

    &:hover {
      color: var(--active);
    }

    &.router-link-exact-active {
      color: var(--active) !important;
    }
  }

  .right {
    width: 40%;
    display: flex;
    justify-content: right;
    align-items: center;

    .login_btn {
      margin-right: 20px;
    }

    .gvb_theme {
      margin-right: 20px;
    }
  }
}


.gvb_nav.show {
  background-color: var(--card_bg);
  box-shadow: 1px 1px 5px #0003;
  color: var(--text);

  a {
    color: var(--text);
  }
}
</style>