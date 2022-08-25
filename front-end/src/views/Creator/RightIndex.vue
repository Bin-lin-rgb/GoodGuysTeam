<template>
  <div class="box content-box">
    <div class="header">
      <span class="title">创作活动<van-icon name="question-o" /></span>
      <span class="more">更多内容 <van-icon name="arrow" /></span>
    </div>
    <div class="banner">
      <el-carousel indicator-position="outside">
        <el-carousel-item v-for="item in 4" :key="item">
          <h3 text="2xl" justify="center">{{ item }}</h3>
        </el-carousel-item>
      </el-carousel>
    </div>

    <div class="user-info">
      {{ userStore.username }}
    </div>
    <div class="box content-box">
      <div class="header">
        <span class="title">内容数据<van-icon name="question-o" /></span>
        <span class="more">更多内容 <van-icon name="arrow" /></span>
      </div>
      <div class="content">
        <div class="content-item">
          <div class="item-title">
            总文章数
            <van-icon name="arrow" style="margin-left: 5px" />
          </div>
          <div class="item-count">{{ state.centerCount.article_cnt }}</div>
        </div>
        <div class="content-item">
          <div class="item-title">
            总专栏数
            <van-icon name="arrow" style="margin-left: 5px" />
          </div>
          <div class="item-count">{{ state.centerCount.column_cnt }}</div>
        </div>
        <div class="content-item">
          <div class="item-title">
            总沸点数
            <van-icon name="arrow" style="margin-left: 5px" />
          </div>
          <div class="item-count">{{ state.centerCount.short_msg_cnt }}</div>
        </div>
        <div class="content-item2">
          <div class="item-title">文章展现数</div>
          <div class="item-count">
            {{ state.cardInfo.all_article_display }}
          </div>
          <div class="item-before">
            较前日<span style="color: #1e80ff">{{
              state.cardInfo.all_article_display
            }}</span>
          </div>
        </div>
        <div class="content-item2">
          <div class="item-title">文章阅读数</div>
          <div class="item-count">{{ state.cardInfo.all_article_view }}</div>
          <div class="item-before">
            较前日<span style="color: #1e80ff">{{
              state.cardInfo.all_article_view
            }}</span>
          </div>
        </div>
        <div class="content-item2">
          <div class="item-title">文章收藏数</div>
          <div class="item-count">
            {{ state.cardInfo.all_article_collect }}
          </div>
          <div class="item-before">
            较前日<span style="color: #1e80ff">{{
              state.cardInfo.all_article_collect
            }}</span>
          </div>
        </div>
      </div>
    </div>
    <div class="box content-box">
      <div class="header">
        <span class="title">关注者数据<van-icon name="question-o" /></span>
        <span class="more">更多内容 <van-icon name="arrow" /></span>
      </div>
      <div class="content">
        <div class="content-item2">
          <div class="item-title">总关注者</div>
          <div class="item-count">{{ state.trend.all_follower }}</div>
          <div class="item-before">
            较前日<span style="color: #1e80ff">{{ state.trend.all_follower }}</span>
          </div>
        </div>
        <div class="content-item2">
          <div class="item-title">活跃关注者</div>
          <div class="item-count">{{ state.trend.incr_active_follower }}</div>
          <div class="item-before">
            较前日<span style="color: #1e80ff">{{
              state.trend.incr_active_follower
            }}</span>
          </div>
        </div>
        <div class="content-item2">
          <div class="item-title">净增关注</div>
          <div class="item-count">{{ state.trend.incr_do_follower }}</div>
          <div class="item-before">
            较前日<span style="color: #1e80ff">{{ state.trend.incr_do_follower }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from "vue";
import { useUserStore } from "@/stores/user";
const userStore = useUserStore();

const state = reactive({
  centerCount: {
    article_cnt: 89,
    column_cnt: 34,
    short_msg_cnt: 780,
  },
  cardInfo: {
    all_article_display: 62,
    all_article_view: 88,
    all_article_collect: 10,
  },
  trend: {
    all_follower: 44,
    incr_active_follower: 9,
    incr_do_follower: 56,
  },
  bannerList: [],
});
</script>

<style lang="less" scoped>
.box {
  background-color: #fff;
  box-sizing: border-box;
  margin-top: 10px;
  padding: 10px;
}
.profile-box {
  display: flex;
  & .avatar {
    width: 60px;
    height: 60px;
  }
  & .user-info {
    flex: 1;
    margin-left: 20px;
    color: #1d2129;
    font-size: 16px;
    font-weight: 500;
  }
  & .home-page {
    width: 80px;
    color: #86909c;
    font-size: 14px;
    font-weight: 400;
  }
}

.content-box {
  & .header {
    position: relative;
    & .title {
      font-weight: 550;
      font-size: 16px;
    }

    & .more {
      color: #86909c;
      font-size: 14px;
      font-weight: 400;
      float: right;
    }
  }

  .banner {
    padding: 15px;
  }

  & .content {
    display: grid;
    grid-template-columns: repeat(auto-fill, 108px);
    justify-content: space-between;
    grid-gap: 15px;
    flex-wrap: wrap;
    margin-top: 10px;
    & .content-item {
      position: relative;
      background-color: #f4f4f4;
      padding: 10px;
      height: 50px;
      & .item-title {
        font-size: 14px;
        color: #86909c;
      }

      & .item-count {
        position: absolute;
        bottom: 10px;
        font-weight: bold;
        font-size: 16px;
      }
    }
    & .content-item2 {
      position: relative;
      padding: 10px;
      border: 1px solid #f4f4f4;
      height: 70px;
      border-radius: 5px;
      & .item-title {
        font-size: 15px;
        color: #86909c;
      }

      & .item-count {
        position: absolute;
        bottom: 30px;
        font-weight: bold;
        font-size: 18px;
      }
      & .item-before {
        font-size: 10px;
        position: absolute;
        bottom: 10px;
      }
    }
  }
}

.el-carousel__item h3 {
  display: flex;
  color: #475669;
  opacity: 0.75;
  line-height: 300px;
  margin: 0;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}
</style>
