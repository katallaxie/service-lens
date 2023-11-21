'use client'

import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger
} from '@/components/ui/dialog'
import { Input } from '@/components/ui/input'
import { useToast } from '@/components/ui/use-toast'

import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { useRouter } from 'next/navigation'
import * as z from 'zod'
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage
} from '@/components/ui/form'
import { Textarea } from '@/components/ui/textarea'
import useSWR from 'swr'

const FormSchema = z.object({
  name: z.string().min(3, {}),
  environment: z.string().min(3, {}),
  description: z
    .string()
    .min(10, {
      message: 'Description must be at least 30 characters.'
    })
    .max(2024, {
      message: 'Description must be less than 2024 characters.'
    }),
  tags: z.array(z.string())
})

const updateWorkloads = (url: string) =>
  fetch('/api/workloads').then(res => res.json())

export function AddProfileDialog() {
  const { data, mutate, isLoading } = useSWR('/api/workloads', updateWorkloads)
  const { toast } = useToast()

  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      name: '',
      description: '',
      environment: '',
      tags: []
    }
  })

  async function onSubmit(form: z.infer<typeof FormSchema>) {
    try {
      await fetch('/api/workloads', {
        method: 'POST',
        body: JSON.stringify(form)
      })

      await mutate({ ...data, rows: [...data.rows, form] }, true)

      toast({
        title: 'Workload created'
      })
    } catch (e) {}
  }

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button variant="outline">Add Profile</Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>New Workload</DialogTitle>
        </DialogHeader>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <div className="grid gap-4 py-4">
              <FormField
                control={form.control}
                name="name"
                render={({ field }) => (
                  <div className="grid w-full max-w-sm items-center gap-1.5">
                    <FormLabel>Name</FormLabel>
                    <FormControl>
                      <Input
                        type="text"
                        disabled={form.formState.isSubmitting}
                        {...field}
                      />
                    </FormControl>
                    <FormDescription>
                      The name of workload to review.
                    </FormDescription>
                    <FormMessage />
                  </div>
                )}
              />
              <FormField
                control={form.control}
                name="description"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Description</FormLabel>
                    <FormControl>
                      <Textarea
                        disabled={form.formState.isSubmitting}
                        placeholder="Tell us a bit about the workload."
                        className="resize-none"
                        {...field}
                      />
                    </FormControl>
                    <FormDescription>
                      The specification to create a lense to review workloads.
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="environment"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Environment</FormLabel>
                    <FormControl>
                      <Input
                        disabled={form.formState.isSubmitting}
                        placeholder="production, staging, development"
                        {...field}
                      />
                    </FormControl>
                    <FormDescription>
                      The environment the workload is running in.
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </div>
            <DialogFooter>
              <Button type="submit" disabled={form.formState.isSubmitting}>
                Create
              </Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  )
}

export default AddProfileDialog