<template>

  <div class="mycard">
    <div class="poker-list">
        <div class="card-item poker bigjoker" style="left: 0px;"></div>
        <div class="card-item poker smalljoker" style="left: 40px;"></div>
        <div class="card-item poker unknown" style="left: 80px;"></div>
    </div>

    <div class="poker-list">
        <div class="card-item poker heart-2" style="left: 0px;"></div>
        <div class="card-item poker heart-3" style="left: 40px;"></div>
        <div class="card-item poker heart-4" style="left: 80px;"></div>
        <div class="card-item poker heart-5" style="left: 120px;"></div>
        <div class="card-item poker heart-6" style="left: 160px;"></div>
        <div class="card-item poker heart-7" style="left: 200px;"></div>
        <div class="card-item poker heart-8" style="left: 240px;"></div>
        <div class="card-item poker heart-9" style="left: 280px;"></div>
        <div class="card-item poker heart-10" style="left: 320px;"></div>
        <div class="card-item poker heart-J" style="left: 360px;"></div>
        <div class="card-item poker heart-Q" style="left: 400px;"></div>
        <div class="card-item poker heart-K" style="left: 440px;"></div>
        <div class="card-item poker heart-A" style="left: 480px;"></div>
    </div>
  </div>

  <div class="ctrl-area">
    <button>开始</button>
    <button>出牌</button>
    <button> 过 </button>
    <button>叫地主</button>
    <button>抢地主</button>
    <button>不叫</button>
  </div>
</template>

<style scoped>
@import '../assets/poker.css';

.mycard {
  margin: auto;
  bottom: 10rem;
  background-color: aliceblue;
  width: 100%;
  min-width: 1000px;
  height: 172px;
}

@media (min-width: 1024px) {
  .mycard {
    margin-top: 0;
    padding: 0.4rem 0 1rem calc(var(--section-gap) / 2);
  }
}
</style>
