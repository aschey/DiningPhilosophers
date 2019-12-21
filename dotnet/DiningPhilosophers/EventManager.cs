using System;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.ComponentModel;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace DiningPhilosophers
{
    static class EventManager
    {
        private static ConcurrentDictionary<string, EventArgs> _eventHandlers = new ConcurrentDictionary<string, EventArgs>();

        public static void Broadcast(string eventName, string arg = null)
        {
            _eventHandlers[eventName].Mutex.WaitOne();
            foreach (var callback in _eventHandlers[eventName].Callbacks)
            {
                new Task(() => callback(arg)).Start();
            }
            if (_eventHandlers[eventName].AutoUnsubscribe)
            {
                _eventHandlers[eventName].Callbacks = new List<Action<string>>();
            }
            _eventHandlers[eventName].Mutex.ReleaseMutex();
        }

        public static void Subscribe(string name, Action<string> callback, bool autoUnsubscribe = true)
        {
            if (_eventHandlers.ContainsKey(name))
            {
                _eventHandlers[name].Mutex.WaitOne();
                _eventHandlers[name].AutoUnsubscribe = autoUnsubscribe;
                _eventHandlers[name].Callbacks.Add(callback);
                _eventHandlers[name].Mutex.ReleaseMutex();
            }
            else
            {
                var args = new EventArgs
                {
                    AutoUnsubscribe = autoUnsubscribe
                };
                args.Callbacks.Add(callback);
                _eventHandlers[name] = args;
            }
        }
    }

    public class EventArgs
    {
        public List<Action<string>> Callbacks = new List<Action<string>>();

        public bool AutoUnsubscribe = true;

        public Mutex Mutex = new Mutex();
    }
}
