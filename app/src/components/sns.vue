<template>
    <div>
      <h3>SNS</h3>
        <div class="link-sns">
        <a class="link" id="linkGitHub" target="_blank" rel="noopener noreferrer" href="https://github.com/ken7python">
          <img class="logo" :src="githubLogo" alt="GitHub Logo" id="icon-github">
          @ken7python
        </a>

        <a class="link" id="linkX" target="_blank" rel="noopener noreferrer" href="https://x.com/ken7python">
          <img class="logo" :src="xLogo" alt="X Logo" id="icon-x">
          @ken7python
        </a>
        <a class="link" id="linkInstagram" target="_blank" rel="noopener noreferrer" href="https://www.instagram.com/ken7python/">
          <img class="logo" :src="InstagramLogo" alt="Instagram Logo" id="icon-instagram">
          @ken7python
        </a>
        <a class="link" id="linkNote" target="_blank" rel="noopener noreferrer" href="https://note.com/ken7python">
          <img class="logo" :src="noteLogo" alt="Note Logo" id="icon-note">
          @ken7python
        </a>
        <a class="link" id="linkYouTube" target="_blank" rel="noopener noreferrer" href="https://www.youtube.com/@KenCodeHanabi">
          <img class="logo" :src="youtubeLogo" alt="YouTube Logo" id="icon-youtube">
          @KenCodeHanabi
        </a>
      </div>
      <hr>
    </div>
</template>

<script setup lang="ts">
  const githubLogo = './github-logo.png'
  const xLogo = './x-logo.png'
  const InstagramLogo = './Instagram-logo.png'
  const noteLogo = './note-logo.png'
  const youtubeLogo = './youtube-logo.png'

</script>

<style scoped>
  .link-sns {
    display: flex;
    flex-direction: column;
    /* align-items: center; */
  }
  .logo {
    height: 18px;
    padding-right: 8px;
    vertical-align: middle;
  }
  .link-sns a {
    color: #333;
    text-decoration: none;
    font-size: 16px;
    margin-bottom: 8px;
    display: flex;
    align-items: center;
  }

  .link{
    padding-bottom: 4px;
  }
</style>