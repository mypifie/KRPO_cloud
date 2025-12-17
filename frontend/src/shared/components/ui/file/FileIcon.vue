<template>
  <div
    @click="() => emit('handleClick')"
    class="max-w-20 flex flex-col items-center justify-center w-full h-full cursor-pointer"
  >
    <FileIcon class="size-12"></FileIcon>
    <Label class="max-w-20 overflow-hidden text-ellipsis">{{fileName}}</Label>
  </div>
</template>

<script setup lang="ts">
import { FileIcon } from 'lucide-vue-next'
import { Label } from '@/shared/components/ui/label'

defineProps<{
  fileName: string
}>();

const emit = defineEmits<{
  (e: 'handleClick'): void
}>()
</script>

<style scoped></style>
